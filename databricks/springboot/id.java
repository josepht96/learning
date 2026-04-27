@EventListener(ApplicationReadyEvent.class)
public void setTaskId() {
    log.info("EcsTaskIdentifier starting...");
    try {
        String metadataUri = System.getenv("ECS_CONTAINER_METADATA_URI_V4");
        log.info("ECS_CONTAINER_METADATA_URI_V4={}", metadataUri);
        if (metadataUri == null) {
            MDC.put("taskId", "local");
            log.info("Running locally, taskId set to 'local'");
            return;
        }
        // ... rest of method
    }
}