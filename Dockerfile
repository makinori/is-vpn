FROM denoland/deno:alpine

WORKDIR /app
COPY app/. /app/.
RUN deno run deps.ts

CMD ["deno", "run", "--allow-net", "--allow-env", "--allow-read", "main.ts"]