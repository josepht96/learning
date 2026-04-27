
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.slf4j.MDC;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.event.EventListener;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

@Component
public class EcsTaskIdentifier {

    private static final Logger log = LoggerFactory.getLogger(EcsTaskIdentifier.class);

    @EventListener(ApplicationReadyEvent.class)
    public void setTaskId() {
        log.info("EcsTaskIdentifier starting...");
        try {
            String metadataUri = System.getenv("ECS_CONTAINER_METADATA_URI_V4");
            log.info("ECS_CONTAINER_METADATA_URI_V4={}", metadataUri);

            if (metadataUri == null) {
                String hostname = java.net.InetAddress.getLocalHost().getHostName();
                MDC.put("taskId", hostname);
                log.info("Running locally, taskId set to hostname: {}", hostname);
                return;
            }

            RestTemplate restTemplate = new RestTemplate();
            String response = restTemplate.getForObject(metadataUri + "/task", String.class);
            JsonNode json = new ObjectMapper().readTree(response);
            String taskArn = json.get("TaskARN").asText();
            String shortId = taskArn.substring(taskArn.lastIndexOf("/") + 1);
            MDC.put("taskId", shortId);
            log.info("Running on ECS, taskId set to: {}", shortId);

        } catch (Exception e) {
            log.warn("Failed to determine taskId, falling back to unknown: {}", e.getMessage());
            MDC.put("taskId", "unknown");
        }
    }
}