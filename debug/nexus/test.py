#!/usr/bin/env python3

import requests
import time
import logging
from datetime import datetime

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(message)s",
    handlers=[
        logging.FileHandler("nexus_latency.log"),
        logging.StreamHandler()
    ]
)

logger = logging.getLogger(__name__)

NEXUS_BASE_URL = "https://your-nexus-instance"
REPO = "ocp-quay"
IMAGE = "openshift-release-dev/ocp-v4.0-art-dev"
AUTH = ("username", "password")  # or None if unauthenticated

def get_tags() -> list:
    url = f"{NEXUS_BASE_URL}/v2/{IMAGE}/tags/list"
    logger.info(f"Fetching tags from {url}")

    start = time.time()
    response = requests.get(url, auth=AUTH, verify=False, timeout=30)
    elapsed = time.time() - start

    logger.info(f"Tags response: {response.status_code} in {elapsed:.3f}s")
    response.raise_for_status()

    tags = response.json().get("tags", [])
    logger.info(f"Found {len(tags)} tags")
    return tags


def get_manifest(tag: str) -> dict:
    url = f"{NEXUS_BASE_URL}/v2/{IMAGE}/manifests/{tag}"
    headers = {
        "Accept": "application/vnd.oci.image.manifest.v1+json,"
                  "application/vnd.docker.distribution.manifest.v2+json,"
                  "application/vnd.docker.distribution.manifest.list.v2+json"
    }

    start = time.time()
    response = requests.get(url, auth=AUTH, headers=headers, verify=False, timeout=30)
    elapsed = time.time() - start

    return {
        "tag": tag,
        "status_code": response.status_code,
        "latency_ms": round(elapsed * 1000, 2),
        "content_length": len(response.content)
    }


def run():
    tags = get_tags()

    results = []
    for i, tag in enumerate(tags):
        logger.info(f"[{i+1}/{len(tags)}] Fetching manifest for tag: {tag}")
        try:
            result = get_manifest(tag)
            results.append(result)
            logger.info(f"  Status: {result['status_code']} | Latency: {result['latency_ms']}ms | Size: {result['content_length']} bytes")
        except requests.exceptions.RequestException as e:
            logger.error(f"  Failed to fetch manifest for {tag}: {e}")

    # Summary stats
    if results:
        latencies = [r["latency_ms"] for r in results if r["status_code"] == 200]
        logger.info("--- Summary ---")
        logger.info(f"Total manifests fetched: {len(results)}")
        logger.info(f"Successful: {sum(1 for r in results if r['status_code'] == 200)}")
        logger.info(f"Failed: {sum(1 for r in results if r['status_code'] != 200)}")
        if latencies:
            logger.info(f"Avg latency: {sum(latencies)/len(latencies):.2f}ms")
            logger.info(f"Min latency: {min(latencies):.2f}ms")
            logger.info(f"Max latency: {max(latencies):.2f}ms")


if __name__ == "__main__":
    run()