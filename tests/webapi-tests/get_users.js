import http from "k6/http";
import { sleep } from "k6";

export let options = {
  stages: [
    { duration: "30s", target: 20 },
    { duration: "30s", target: 50 },
    { duration: "30s", target: 100 },
    { duration: "30s", target: 0 },
  ],
};

export default function () {
    const iter = __ITER + 1;
    const email = `email${iter}@box.ru`;

    const url = `http://localhost:8080/api/user?email=${encodeURIComponent(email)}`;
    http.get(url);
}

export function handleSummary(data) {
    const metrics = data.metrics;

    const summaryText = `
==================== K6 GET Test Summary =====================
Total requests:         ${metrics.http_reqs.count}
Successful requests:    ${metrics.http_reqs.passes}
Failed requests:        ${metrics.http_reqs.fails}

Response Time (ms):
  Avg:                  ${metrics.http_req_duration.avg.toFixed(2)}
  Min:                  ${metrics.http_req_duration.min.toFixed(2)}
  Max:                  ${metrics.http_req_duration.max.toFixed(2)}
  90th percentile:      ${metrics.http_req_duration['p(90)']}
  95th percentile:      ${metrics.http_req_duration['p(95)']}

==============================================================
`;

    return {
        "get_summary.json": JSON.stringify(data, null, 2),
        "get_summary.txt": summaryText
    };
}
