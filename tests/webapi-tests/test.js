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
    const payload = JSON.stringify({
        password: `user_${Math.random()}`,
        email: `email_${Math.random()}@test.com`,
    });


    const headers = {
        "Content-Type": "application/json",
    };

    http.post("http://localhost:8080/api/user/register", payload, { headers });
   
}

export function handleSummary(data) {
    // Извлекаем нужные метрики
    const metrics = data.metrics;

    // Создаем кастомный текстовый отчет
    const summaryText = `
    ==== K6 Test Summary ====
    Total requests: ${metrics.http_reqs.count}
    Successful requests: ${metrics.http_reqs.passes}
    Failed requests: ${metrics.http_reqs.fails}

    Response Time (ms):
      Avg: ${metrics.http_req_duration.avg.toFixed(2)}
      Min: ${metrics.http_req_duration.min.toFixed(2)}
      Max: ${metrics.http_req_duration.max.toFixed(2)}
      90th percentile: ${metrics.http_req_duration['p(90)']}
      95th percentile: ${metrics.http_req_duration['p(95)']}

    ================================================
    `;

    // Возвращаем объект с файлами, которые нужно сохранить
    return {
        "summary.json": JSON.stringify(data, null, 2),  // полный JSON с метриками
        "summary.txt": summaryText                      // человекочитаемый отчет
    };
}