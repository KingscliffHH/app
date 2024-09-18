### BUILD ###
FROM node:20-alpine as BUILD

### DEFAULT ENV VARWS ###
ARG VITE_API_BASE_URL=/api
ARG VERSION=0.0.0

# Copy all the required build files
WORKDIR /app
COPY package.json /app/
COPY pnpm-lock.yaml /app/

# Install the dependencies
RUN set -xe \
    && npm set progress=false \
    && npm install pnpm --location=global \
    && pnpm set progress=false \
    && pnpm audit --audit-level=high \
    && pnpm install

COPY . /app/

# Perform build
RUN pnpm version ${VERSION} --no-git-tag-version --allow-same-version \
    && pnpm build

### IMAGE BUILD STAGE ###
FROM nginx:1.17-alpine

# Redirect missing files (404) to Angular (required for Angular router to work)
# and purge the existing default application file
RUN sed -i  '/location \/ {/ a \ \ \ \ \ \ \ try_files $uri $uri/ /index.html?$args;' /etc/nginx/conf.d/default.conf \
    && rm -rf /usr/share/nginx/html/*

# Copy the UI application
COPY  --from=BUILD /app/dist/ /usr/share/nginx/html
COPY ./entrypoint.sh ./entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh" ]
CMD ["nginx", "-g", "daemon off;"]

