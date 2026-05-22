# GraphQL — Complete Beginner-to-Intermediate Notes
## With Elmeasure Examples

---

# 1. What is GraphQL?

:contentReference[oaicite:0]{index=0} is a query language and runtime for APIs.

It allows the frontend to request exactly the data it needs from the backend.

Created by:
:contentReference[oaicite:1]{index=1}

---

# 2. Why GraphQL Was Created

Traditional REST APIs caused problems:

- Too many API endpoints
- Multiple network requests
- Overfetching unnecessary data
- Frontend/backend coordination complexity

Example:

```txt
GET /plants/1
GET /plants/1/meters
GET /meters/11/alerts
GET /meters/11/readings
```

Frontend may need many requests for one screen.

GraphQL solves this by allowing one structured query.

---

# 3. Core Idea of GraphQL

Frontend specifies:
- what data it wants
- which fields it needs
- how deeply nested relationships should go

Backend returns exactly that structure.

---

# 4. REST vs GraphQL

| REST | GraphQL |
|---|---|
| Many endpoints | Single endpoint |
| Backend decides response shape | Frontend decides |
| Overfetching common | Exact fields fetched |
| Easier initially | More flexible |
| Simple caching | Harder caching |
| Good for simple APIs | Good for connected systems |

---

# 5. Example — REST vs GraphQL

## REST

```txt
GET /plant/1
GET /plant/1/meters
GET /meter/11/latest
GET /meter/11/alerts
```

---

## GraphQL

```graphql
query {
  plant(id: 1) {
    name

    meters {
      name

      latestReading {
        value
      }

      alerts {
        severity
      }
    }
  }
}
```

Single request.

---

# 6. Why It Is Called GraphQL

Data in real systems is connected:

```txt
Company
 └── Plant
      └── Meter
           └── Reading
                └── Alert
```

This connected structure behaves like a graph.

GraphQL allows traversal through these relationships.

---

# 7. Main Components of GraphQL

---

## 7.1 Schema

Defines:
- available data
- relationships
- data types

Example:

```graphql
type Plant {
  id: ID!
  name: String!
  meters: [Meter]
}

type Meter {
  id: ID!
  name: String!
  latestReading: Float
}
```

Schema acts like:
- API contract
- blueprint
- type system

---

## 7.2 Query

Used to fetch data.

Equivalent to:
```txt
GET
```

Example:

```graphql
query {
  plant(id: 1) {
    name
  }
}
```

---

## 7.3 Mutation

Used to modify data.

Equivalent to:
```txt
POST
PUT
DELETE
```

Example:

```graphql
mutation {
  createMeter(name: "Boiler Meter") {
    id
    name
  }
}
```

---

## 7.4 Resolver

Resolver is backend logic that fetches data.

Example:

```txt
plant → meters
```

Resolver may:
- query database
- call another service
- perform calculations

Resolvers power GraphQL internally.

---

# 8. How GraphQL Works Internally

```txt
Frontend
    ↓
GraphQL Query
    ↓
GraphQL Server
    ↓
Schema Validation
    ↓
Resolvers Execute
    ↓
Database / Services
    ↓
JSON Response Returned
```

---

# 9. GraphQL Response Structure

Example:

```json
{
  "data": {
    "plant": {
      "name": "Bangalore Plant"
    }
  }
}
```

GraphQL responses are usually tree-shaped JSON.

---

# 10. Context in GraphQL

One of GraphQL’s biggest strengths.

Example:

```graphql
query {
  plant(id: 1) {
    name

    meters {
      name

      latestReading {
        value
      }
    }
  }
}
```

The reading automatically comes with context:
- which plant
- which meter
- hierarchy relationships

Frontend receives connected information instead of isolated records.

---

# 11. Overfetching and Underfetching

---

## Overfetching

Fetching unnecessary data.

Example:

```json
{
  "id": 11,
  "name": "Meter",
  "serialNo": "ABC",
  "manufacturer": "XYZ",
  "latestReading": 420
}
```

Frontend only needed:
```txt
name + latestReading
```

---

## Underfetching

Not getting enough data in one request.

Requires multiple API calls.

GraphQL solves both problems.

---

# 12. Why Frontend Developers Like GraphQL

Frontend controls:
- fields
- nesting
- response structure

Backend no longer needs:
- custom dashboard endpoints
- many specialized APIs

Frontend becomes more flexible.

---

# 13. GraphQL and Databases

GraphQL is NOT a database.

It sits between:
- frontend
- backend/data sources

Example architecture:

```txt
Frontend
    ↓
GraphQL API Layer
    ↓
Database / Services
```

---

# 14. Best Databases with GraphQL

Good combinations:

- :contentReference[oaicite:2]{index=2}
- :contentReference[oaicite:3]{index=3}
- :contentReference[oaicite:4]{index=4}

---

# 15. GraphQL in Elmeasure

Elmeasure data is naturally hierarchical:

```txt
Organization
 └── Plant
      └── Building
           └── Meter
                └── Reading
                     └── Alert
```

GraphQL is excellent for:
- dashboards
- enterprise hierarchy browsing
- reports
- customer portals
- analytics summaries

---

# 16. Good Use Cases in Elmeasure

---

## Dashboard APIs

Example:

```graphql
query {
  plant(id: 1) {
    energyToday

    meters {
      latestReading
    }
  }
}
```

---

## Multi-tenant Enterprise Systems

Useful when:
- many organizations
- many plants
- deeply connected data

---

## Custom Dashboards

Different clients need different metrics.

GraphQL lets frontend request only relevant fields.

---

## Mobile Apps

Reduces unnecessary payload size.

---

## API Gateway for Microservices

GraphQL can unify:
- EMS service
- Alert service
- Analytics service
- Billing service

into one API.

---

# 17. Where GraphQL is BAD

---

## High-Speed IoT Ingestion

Bad fit for:
- millions of readings
- telemetry streams
- real-time sensor ingestion

Use:
- :contentReference[oaicite:5]{index=5}
- :contentReference[oaicite:6]{index=6}
- REST ingestion APIs
- gRPC

---

## Heavy Historical Analytics

Example:

```txt
5 years of second-by-second readings
```

Can become inefficient.

Better:
- analytics engines
- batch pipelines
- pre-aggregations

---

## Tiny Simple Applications

REST is simpler for:
- login
- health checks
- basic CRUD

---

# 18. Recommended Elmeasure Architecture

```txt
IoT Devices
    ↓
MQTT / Kafka
    ↓
Ingestion Services
    ↓
TimescaleDB / PostgreSQL
    ↓
Business Logic Services
    ↓
GraphQL API Layer
    ↓
Frontend Dashboard
```

---

# 19. Important GraphQL Concepts

---

## Strong Typing

GraphQL schemas are typed.

Example:

```graphql
type Meter {
  latestReading: Float
}
```

Benefits:
- safer APIs
- better developer tooling
- autocomplete
- validation

---

## Nested Queries

GraphQL naturally supports hierarchy.

Example:

```graphql
query {
  company(id: 1) {
    plants {
      meters {
        alerts {
          severity
        }
      }
    }
  }
}
```

---

## Single Endpoint

Usually:

```txt
POST /graphql
```

Unlike REST:
- many endpoints not needed

---

# 20. GraphQL Does NOT Automatically Do AI

GraphQL only follows relationships you define.

It does NOT infer meaning.

It cannot automatically conclude:

```txt
High energy usage means motor inefficiency
```

That requires:
- analytics
- AI
- business logic

---

# 21. GraphQL Security Concerns

GraphQL can be dangerous if unmanaged.

---

## Deep Query Attacks

Client may request massive nested data.

Example:

```txt
company → plants → meters → readings → alerts
```

Need:
- depth limits
- query cost limits

---

## Expensive Queries

Large nested queries can overload backend.

Need:
- optimization
- pagination
- caching

---

# 22. Pagination

Important in GraphQL.

Example:

```graphql
query {
  meters(limit: 10, offset: 0) {
    name
  }
}
```

Prevents huge responses.

---

# 23. Subscriptions (Real-Time GraphQL)

GraphQL supports real-time updates.

Example:
- live meter readings
- alert notifications

Uses:
- WebSockets

Example:

```graphql
subscription {
  meterUpdated(id: 11) {
    latestReading
  }
}
```

Useful for dashboards.

---

# 24. Popular GraphQL Tools

---

## Backend

- :contentReference[oaicite:7]{index=7}
- :contentReference[oaicite:8]{index=8}
- :contentReference[oaicite:9]{index=9}

---

## Frontend

- :contentReference[oaicite:10]{index=10}
- :contentReference[oaicite:11]{index=11}

---

# 25. Best Learning Path

---

## Step 1

Learn:
- APIs
- JSON
- REST

---

## Step 2

Understand:
- schema
- query
- mutation
- resolver

---

## Step 3

Build simple GraphQL API:
- users
- posts
- comments

---

## Step 4

Add:
- authentication
- pagination
- subscriptions

---

## Step 5

Learn optimization:
- batching
- caching
- query complexity control

---

# 26. Most Important Concept to Remember

GraphQL is not mainly about:
- fewer endpoints
- one request

The REAL value is:

# GraphQL models connected data naturally.

It mirrors real-world relationships cleanly and allows frontend applications to traverse those relationships flexibly.

---

# 27. Final Recommendation for Elmeasure

Use GraphQL for:
- dashboards
- enterprise hierarchy browsing
- reports
- mobile apps
- customer portals
- aggregated APIs

Do NOT use GraphQL for:
- telemetry ingestion
- streaming pipelines
- raw industrial event processing

Best architecture:

```txt
Kafka/MQTT + Time-Series Database + GraphQL API Layer
```

This is modern, scalable, and ideal for industrial SaaS platforms like Elmeasure.