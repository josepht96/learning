package gov.fda.cber.parser;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.web.client.RestTemplate;

public class ContainerIdentifier {

    private static String CACHED_ID = null;

    public static String getTaskId() {
        if (CACHED_ID != null) {
            return CACHED_ID;
        }
        try {
            String metadataUri = System.getenv("ECS_CONTAINER_METADATA_URI_V4");
            if (metadataUri == null) {
                // Local — use hostname
                CACHED_ID = java.net.InetAddress.getLocalHost().getHostName();
            } else {
                // ECS — use short task ARN
                RestTemplate restTemplate = new RestTemplate();
                String response = restTemplate.getForObject(metadataUri + "/task", String.class);
                JsonNode json = new ObjectMapper().readTree(response);
                String taskArn = json.get("TaskARN").asText();
                CACHED_ID = taskArn.substring(taskArn.lastIndexOf("/") + 1);
            }
        } catch (Exception e) {
            CACHED_ID = "unknown";
        }
        return CACHED_ID;
    }
}