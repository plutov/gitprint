FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci
COPY . .

ARG VITE_API_ADDR
ARG VITE_ANALYTICS_ENABLED

ENV VITE_API_ADDR=${VITE_API_ADDR}
ENV VITE_ANALYTICS_ENABLED=${VITE_ANALYTICS_ENABLED}

RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 5173

CMD ["nginx", "-g", "daemon off;"]
