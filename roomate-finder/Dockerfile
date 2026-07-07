FROM node:20-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# 1. Build the app during the image creation phase
RUN npm run build

# 2. Expose the port Next.js runs on
EXPOSE 3000

# 3. Start the actual live production server
CMD ["npm", "run", "start"]
