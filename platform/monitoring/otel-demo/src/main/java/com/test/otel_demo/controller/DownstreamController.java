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
@RequestMapping("/downstream")
public class DownstreamController {

    private static final Logger logger = LoggerFactory.getLogger(DownstreamController.class);
    private final RestTemplate restTemplate;
    
    @Value("${downstream.service.url:http://localhost:9000}")  // Default to localhost:9000 if not set
    private String downstreamServiceUrl;

    public DownstreamController(RestTemplate restTemplate) {
        this.restTemplate = restTemplate;
    }

    @GetMapping
    public Map<String, Object> callDownstreamService() {
        logger.info("Calling downstream service at {}", downstreamServiceUrl);
        try {
            Map response = restTemplate.getForObject(downstreamServiceUrl, Map.class);
            logger.info("Downstream response: {}", response);
            return response;
        } catch (Exception e) {
            logger.error("Failed to reach downstream service", e);
            return Map.of("error", "Downstream service unreachable");
        }
    }
}
