public private static String CACHED_TASK_ID = null;

private String taskId() {
    if (CACHED_TASK_ID != null) {
        return CACHED_TASK_ID;
    }
    try {
        String metadataUri = System.getenv("ECS_CONTAINER_METADATA_URI_V4");
        if (metadataUri == null) {
            CACHED_TASK_ID = java.net.InetAddress.getLocalHost().getHostName();
        } else {
            RestTemplate restTemplate = new RestTemplate();
            String response = restTemplate.getForObject(metadataUri + "/task", String.class);
            JsonNode json = new ObjectMapper().readTree(response);
            String taskArn = json.get("TaskARN").asText();
            CACHED_TASK_ID = taskArn.substring(taskArn.lastIndexOf("/") + 1);
        }
    } catch (Exception e) {
        CACHED_TASK_ID = "unknown";
    }
    return CACHED_TASK_ID;
} {
    
}
