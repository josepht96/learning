The other root is Sectigo's **AAA Certificate Services** (p11-kit lists it as "Comodo AAA Services root"). Removing the eMudhra root is not enough wherever Sectigo's AAA Certificate Services root is still trusted, because the same eight keys were cross-signed under it.

For UBI8, I'd do three things in the Dockerfile: pull the patched package, explicitly blocklist both roots, and fail the build if either is still trusted.

**1. Update `ca-certificates`.** Red Hat already shipped a fix: it's tracked as CVE-2026-114406, with RHSAs for RHEL 7.9 through 10.2, and Red Hat says to rebuild containers built on UBI. I'd still add the explicit blocklist below. It's idempotent, it covers you if a mirror or Satellite lags behind, and I couldn't confirm that Red Hat's errata also drops the AAA root.

**2. Blocklist both roots with p11-kit.** This is the standard RHEL approach. Don't delete PEMs from the extracted bundle, because the next `update-ca-trust` would regenerate them.

```dockerfile
FROM registry.access.redhat.com/ubi8/ubi:8.10
USER 0

# 1. Pull the patched trust store (CVE-2026-114406)
RUN dnf -y update ca-certificates && dnf clean all

# 2. Explicitly distrust the two roots (no-op if the package already dropped them)
RUN set -e; \
    d=/etc/pki/ca-trust/source/blocklist; [ -d "$d" ] || d=/etc/pki/ca-trust/source/blacklist; \
    for obj in "emSign%20ECC%20Root%20CA%20-%20G3" "Comodo%20AAA%20Services%20root"; do \
      f="$d/$(echo "$obj" | sed 's/%20/_/g').pem"; \
      trust extract --overwrite --format=pem-bundle \
        --filter="pkcs11:object=${obj};type=cert" "$f" \
        || echo "$obj not present in source store; skipping"; \
    done; \
    update-ca-trust extract

# 3. Fail the build if either root is still trusted (PEM bundle + Java keystore)
RUN set -e; \
    if openssl storeutl -noout -text -certs /etc/pki/tls/certs/ca-bundle.crt \
       | grep -E "CN ?= ?(emSign ECC Root CA - G3|AAA Certificate Services)"; then \
      echo "Distrusted root still present" >&2; exit 1; fi; \
    if command -v keytool >/dev/null; then \
      ! keytool -list -keystore /etc/pki/ca-trust/extracted/java/cacerts -storepass changeit \
        | grep -Ei "emsign ecc root ca - g3|aaa certificate services"; fi

# USER 1001  # restore your non-root user if applicable
```

The blocklist snippet follows the RHEL workaround published in Konfirmity's advisory. I adapted it to skip cleanly when the package has already removed a cert. On `ubi8/ubi-minimal`, swap `dnf` for `microdnf`. `trust` comes from `p11-kit-trust`, which is a dependency of `ca-certificates`, so it should already be there.

**3. Check runtimes that ignore the system store.** On RHEL, `update-ca-trust` also regenerates `/etc/pki/ca-trust/extracted/java/cacerts`. Red Hat's OpenJDK packages, including the `ubi8/openjdk-*` images, point at that file, so Spring Boot apps on those images are covered. Anything that carries its own trust store needs separate attention:

- **A JDK from a tarball (Temurin, Zulu, etc.):** it has its own `$JAVA_HOME/lib/security/cacerts`. Either point it at the system store or update the JDK. OpenJDK's second fix (CVE-2026-114820, Oct 24) removes the emSign G1, G2 and ECC G3 roots and Sectigo's AAA root.
- **Go:** fixed in 1.26.10 / 1.27.3, and the fix takes effect only when a program is rebuilt.
- **Python and Node:** certifi 2026.10.19 or later; Node 22.23.4, 24.22.1 or 26.12.1.

**Two judgment calls:**

- **Removing the AAA root can break old chains.** In Konfirmity's tests, removing AAA Certificate Services did not break connections to current Sectigo-issued sites, because their chains also lead to the USERTrust roots. Older clients that lack those roots could still be affected. Test against any internal or partner endpoints with older Sectigo chains before rolling it out.
- **You may want to go further than G3.** Chrome says it is investigating the continued inclusion of eMudhra's remaining CA certificates, which includes the G1 root. OpenJDK already removes it. If you want the base image to match the strictest posture, add `emSign%20Root%20CA%20-%20G1` (and G2) to the loop.

**Live check after building.** Run this from inside the container:

```bash
curl -sS -o /dev/null https://dv-g3b-emudhra.we-have-a-problem.com/;     echo $?   # want 60
curl -sS -o /dev/null https://dv-g3b-aaa-emudhra.we-have-a-problem.com/; echo $?   # want 60
```

Exit 60 means the certificate was rejected, which is what you want. Behind TLS inspection (an egress proxy, Zscaler, etc.), these checks can report a fixed machine as still trusting the chain, because the network answers with its own certificate. On a federal network, confirm the issuer with `openssl s_client` before trusting the result.

Sources:
- [Konfirmity: eMudhra Intermediate CA Compromise Advisory](https://www.konfirmity.com/advisories/emudhra-intermediate-ca-compromise)
- [Red Hat CVE-2026-114406](https://access.redhat.com/security/cve/CVE-2026-114406)
- [OpenJDK eMudhra guidance](https://openjdk.org/emudhra)
- [Mozilla Bugzilla 2083449](https://bugzilla.mozilla.org/show_bug.cgi?id=2083449)