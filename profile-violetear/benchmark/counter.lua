--  wrk -c2 -d1s -t1 -s counter.lua http://0:8080

counter = 0

request = function()
    path = "/" .. counter
    wrk.headers["X-Counter"] = counter
    counter = counter + 1
    return wrk.format(nil, path)
end


done = function(summary, latency, requests)
    -- open output file
    errors = summary.errors
    failed = errors.connect + errors.read + errors.write + errors.timeout
    -- latency and duration are measured in microseconds
    --for _, p in pairs({50, 66, 75, 80, 90, 95, 98, 99}) do
    --n = latency:percentile(p)
    --f:write(string.format("\"p%dLatencyMs\": %g,", p, n / 1000.0))
    --end

    f = io.open("results/wrk.csv", "a+")
    f:write("date,failedRequests,timeoutRequests,non2xxResponses,maxLatencyMs,avgLatencyMs,completedRequests,requestsPerSecond,kBytesPerSec\n")
    f:write(string.format("%s,%d,%d,%d,%g,%g,%g,%07.2f,%07.2f\n",
    os.date("!%Y-%m-%dT%TZ"),
    failed,
    errors.timeout,
    errors.status,
    latency.max / 1000.0,
    latency.mean / 1000.0,
    summary.requests,
    summary.requests / summary.duration * 1000000,
    summary.bytes / 1024 / summary.duration * 1000000))
    f:close()
end
