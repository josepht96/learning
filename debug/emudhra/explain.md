The keys behind certain eMudhra certificate authorities are publicly recoverable, so anyone can use them to generate a spoofed certificate for any website or internal server. Any container that still trusts these certificate authorities will accept the spoofed certificate as legitimate.

To exploit this, an attacker would need the following:

1. A position on the network path. The container’s traffic has to reach an attacker-controlled host, for example through altered DNS records, ARP spoofing, a compromised proxy, or tampered proxy or hosts settings.
2. A client trusting the affected roots. The workload’s trust store still includes G3 or AAA. The listed software below is the client, in this case.
3. No revocation checking. Containers don't automatically revoke certificates. The containers must be rebuilt to revoke trust for the root CA

Once all three are in place, the attacker can intercept TLS traffic.

In our largely air-gapped environment, the attacker would already need to be inside the network as an insider, a compromised host, or a tampered artifact. The risk is low, and full remediation requires rebuilding all images with the updated trust store.