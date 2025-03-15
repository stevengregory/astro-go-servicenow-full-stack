FROM node:lts-alpine AS runtime

WORKDIR /app

ARG API_URL=http://localhost:8080

COPY . .

RUN npm install
RUN npm run build

ENV PUBLIC_API_URL=$API_URL
ENV HOST=0.0.0.0
ENV PORT=4321
ENV MODE=production

EXPOSE 4321

CMD ["npm", "run", "start"]
