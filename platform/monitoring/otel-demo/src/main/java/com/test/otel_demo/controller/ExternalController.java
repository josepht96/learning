package com.test.otel_demo.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import java.util.Map;

@RestController
@RequestMapping("/external")
public class ExternalController {

    private static final Logger logger = LoggerFactory.getLogger(ExternalController.class);
    private final RestTemplate restTemplate;
    
    @Value("${external.service.url:http://localhost:9000}")  // Default to localhost:9000 if not set
    private String externalServiceUrl;

    public ExternalController(RestTemplate restTemplate) {
        this.restTemplate = restTemplate;
    }

    @GetMapping
    public Map<String, Object> callExternalService() {
        logger.info("Calling external service at {}", externalServiceUrl);
        try {
            Map response = restTemplate.getForObject(externalServiceUrl, Map.class);
            logger.info("External service response: {}", response);
            return response;
        } catch (Exception e) {
            logger.error("Failed to reach external service", e);
            return Map.of("error", "External service unreachable");
        }
    }
}
