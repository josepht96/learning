
import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Gauge;
import io.micrometer.core.instrument.Meter;
import io.micrometer.core.instrument.MeterRegistry;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.actuate.health.HealthComponent;
import org.springframework.boot.actuate.health.HealthEndpoint;
import org.springframework.boot.actuate.health.SystemHealth;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import org.slf4j.MDC;

import org.w3c.dom.css.Counter;

@Component
@EnableScheduling
public class HealthMetricsLogger {

    private static final Logger log = LoggerFactory.getLogger(HealthMetricsLogger.class);

    @Autowired
    private HealthEndpoint healthEndpoint;

    @Autowired
    private MeterRegistry meterRegistry;

    @Scheduled(fixedRate = 30000)
    public void logMetrics() {
        logHealth();
        logMemory();
        logThreads();
        logHttpMetrics();
        logCpu();
    }

    private String taskId() {
        String id = MDC.get("taskId");
        return id != null ? id : "unknown";
    }

    private void logHealth() {
        HealthComponent health = healthEndpoint.health();
        StringBuilder components = new StringBuilder();
        if (health instanceof SystemHealth) {
            ((SystemHealth) health).getComponents().forEach(
                    (name, component) -> components.append(name).append("=").append(component.getStatus()).append(" "));
        }
        log.info("[{}] HEALTH status={} components=[{}]", taskId(), health.getStatus(), components.toString().trim());
    }

    private void logMemory() {
        log.info("[{}] MEMORY heap={}MB heap_committed={}MB heap_max={}MB non_heap={}MB",
                taskId(),
                mbGauge("jvm.memory.used", "area", "heap"),
                mbGauge("jvm.memory.committed", "area", "heap"),
                mbGauge("jvm.memory.max", "area", "heap"),
                mbGauge("jvm.memory.used", "area", "nonheap"));
    }

    private void logThreads() {
        log.info("[{}] THREADS live={} daemon={} peak={} blocked={} waiting={} runnable={}",
                taskId(),
                longGauge("jvm.threads.live"),
                longGauge("jvm.threads.daemon"),
                longGauge("jvm.threads.peak"),
                longGauge("jvm.threads.states", "state", "blocked"),
                longGauge("jvm.threads.states", "state", "waiting"),
                longGauge("jvm.threads.states", "state", "runnable"));
    }

    private void logHttpMetrics() {
        StringBuilder sb = new StringBuilder();
        try {
            meterRegistry.find("http.server.requests").meters().forEach(meter -> {
                String uri = getTag(meter, "uri");
                String method = getTag(meter, "method");
                String status = getTag(meter, "status");
                double count = meter.measure().iterator().next().getValue();
                sb.append(String.format("[%s %s %s count=%d] ", method, uri, status, (long) count));
            });
        } catch (Exception e) {
            sb.append("unavailable");
        }
        log.info("[{}] HTTP {}", taskId(), sb.toString().trim());
    }

    private void logCpu() {
        log.info("[{}] CPU process={}% system={}% load_avg={}",
                taskId(),
                pctGauge("process.cpu.usage"),
                pctGauge("system.cpu.usage"),
                String.format("%.2f", gaugeValue("system.load.average.1m")));
    }

    // --- Helpers ---

    private String mbGauge(String metric, String... tags) {
        double val = gaugeValue(metric, tags);
        return String.format("%.1f", val / 1024 / 1024);
    }

    private long longGauge(String metric, String... tags) {
        return (long) gaugeValue(metric, tags);
    }

    private String pctGauge(String metric, String... tags) {
        return String.format("%.1f", gaugeValue(metric, tags) * 100);
    }

    private double gaugeValue(String metric, String... tags) {
        try {
            Gauge gauge = meterRegistry.find(metric).tags(tags).gauge();
            return gauge != null ? gauge.value() : 0.0;
        } catch (Exception e) {
            log.warn("Metric [{}] unavailable: {}", metric, e.getMessage());
            return 0.0;
        }
    }

    private void logCounter(String metricName, String label) {
        try {
            Counter counter = meterRegistry.find(metricName).counter();
            if (counter != null) {
                log.info("  {}: {}", label, (long) counter.count());
            }
        } catch (Exception e) {
            log.warn("Metric [{}] unavailable: {}", metricName, e.getMessage());
        }
    }

    private String getTag(Meter meter, String tagKey) {
        return meter.getId().getTag(tagKey) != null ? meter.getId().getTag(tagKey) : "unknown";
    }
}