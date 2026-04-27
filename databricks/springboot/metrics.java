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

    private void logHealth() {
        HealthComponent health = healthEndpoint.health();
        log.info("=== Health ===");
        log.info("Status: {}", health.getStatus());
        if (health instanceof SystemHealth) {
            ((SystemHealth) health).getComponents().forEach((name, component) ->
                log.info("  Component [{}]: {}", name, component.getStatus())
            );
        }
    }

    private void logMemory() {
        log.info("=== Memory ===");
        // Heap
        logGauge("jvm.memory.used",      "heap",      "area", "heap");
        logGauge("jvm.memory.committed", "heap committed", "area", "heap");
        logGauge("jvm.memory.max",       "heap max",  "area", "heap");
        // Non-heap
        logGauge("jvm.memory.used",      "non-heap",  "area", "nonheap");
        // GC pressure
        logCounter("jvm.gc.pause", "GC pause count");
    }

    private void logThreads() {
        log.info("=== Threads ===");
        logGauge("jvm.threads.live",   "live");
        logGauge("jvm.threads.daemon", "daemon");
        logGauge("jvm.threads.peak",   "peak");
        logGauge("jvm.threads.states", "blocked",  "state", "blocked");
        logGauge("jvm.threads.states", "waiting",  "state", "waiting");
        logGauge("jvm.threads.states", "runnable", "state", "runnable");
    }

    private void logHttpMetrics() {
        log.info("=== HTTP ===");
        try {
            meterRegistry.find("http.server.requests").meters().forEach(meter -> {
                String uri    = getTag(meter, "uri");
                String method = getTag(meter, "method");
                String status = getTag(meter, "status");
                double count  = meter.measure().iterator().next().getValue();
                log.info("  [{} {}] status={} count={}", method, uri, status, (long) count);
            });
        } catch (Exception e) {
            log.warn("HTTP metrics unavailable: {}", e.getMessage());
        }
    }

    private void logCpu() {
        log.info("=== CPU ===");
        logGauge("process.cpu.usage",  "process");
        logGauge("system.cpu.usage",   "system");
        logGauge("system.load.average.1m", "load avg 1m");
    }

    // --- Helpers ---

    private void logGauge(String metricName, String label, String... tags) {
        try {
            Gauge gauge = meterRegistry.find(metricName).tags(tags).gauge();
            if (gauge != null) {
                double value = gauge.value();
                if (metricName.contains("memory")) {
                    log.info("  {}: {} MB", label, String.format("%.1f", value / 1024 / 1024));
                } else if (metricName.contains("cpu") || metricName.contains("usage")) {
                    log.info("  {}: {}%", label, String.format("%.1f", value * 100));
                } else {
                    log.info("  {}: {}", label, (long) value);
                }
            }
        } catch (Exception e) {
            log.warn("Metric [{}] unavailable: {}", metricName, e.getMessage());
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