package com.test.otel_demo.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Map;

@RestController
@RequestMapping("/")
public class HelloController {
    private static final Logger logger = LoggerFactory.getLogger(HelloController.class);
    
    @GetMapping
    public Map<String, String> helloWorld() {
        logger.info("Received request at '/' endpoint");
        return Map.of("message", "hello, world");
    }
}