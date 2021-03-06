# Microservices

This is demo / sample / example project using microservices architecture for Online Food Delivery App.

# Install

**Step1** - clone this repo
```
git clone https://github.com/nurali-techie/microservices.git
```

**Step2** - run setup.sh to setup all infra, services using docker-compose
```
./setup.sh
```

Use `docker-compose down` to remove all.

# Architecture

## Services
- menu-service - manages restaurant and menu entity
- search-service - manages search query
- customer-webui - customer webui frontend

## Infra
- kong - API Gateway
- kafka - async communication between services

![components](components.png)

## Flows
- create restaurant using menu_service
- create menu items using menu_service
- open customer webui home page
- search menu item from search page

![flows](flows.png)

# References

- REST API Design [link](https://www.mscharhag.com/p/rest-api-design)
- Stripe API Docs [link](https://stripe.com/docs/api)
- Code: kong docker compose [link](https://github.com/Kong/demo-scene/tree/main/kong-docker)
- Video: kong docker compose setup [link](https://youtu.be/sJEID1xEZMg)
- Kafka topic naming [link](https://www.xeotek.com/topic-naming-conventions-how-do-i-name-my-topics-5-recommendations-with-examples/)
- Video: Event-driven Architectures on Apache Kafka with Spring Boot [link](https://youtu.be/xyaFygU9C2Q)
- Video: Consuming REST Web Service in an HTML Page [link](https://youtu.be/KjNXOi4Wqbk)
- Getting Started with Redis and Go - Tutorial [link](https://tutorialedge.net/golang/go-redis-tutorial/)
