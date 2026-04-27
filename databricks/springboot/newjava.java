private static String CACHED_TASK_ID = null;

private String taskId() {
    if (CACHED_TASK_ID != null) {
        return CACHED_TASK_ID;
    }
    try {
        CACHED_TASK_ID = java.net.InetAddress.getLocalHost().getHostName();
    } catch (Exception e) {
        CACHED_TASK_ID = "unknown";
    }
    return CACHED_TASK_ID;
}