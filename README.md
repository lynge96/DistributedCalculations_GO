# DistributedCalculations_GO

### **GO omskrivning**
Har omskrevet den forrige løsning skrevet i .NET til en GO-løsning.
Alt logik fungerer mere eller mindre som den forrige version.

Derudover har jeg udvidet løsningen med mere funktionalitet ud fra nice to have kravene.
Jeg har først og fremmest fået implementeret en reverse proxy/Nginx som er vigtig i en microservice arkitektur, som vi snakkede om til samtalen.
Derudover har jeg udarbejdet en Authenticator API, hvor man kan oprette sig selv som en bruger, samt login og logout.
Bruger JWT token til authentication af endpoints for Calculator og CalculatorHistory API'erne.

### **Systemarkitektur**
![Systemdiagram.png](assets/Systemdiagram.png)

## Services

**Calculator API** – modtager matematiske udtryk, evaluerer dem og publicerer
resultater til RabbitMQ.

**Calculation History API** – consumer beregningsresultater fra RabbitMQ og
gemmer de seneste 5 i hukommelsen.

**Authenticator API** – håndterer brugeroprettelse, login og logout.
Udsteder JWT tokens til brug på beskyttede endpoints.

## Endpoints

| Method   | Endpoint                | Beskrivelse                    | Auth |
|----------|-------------------------|--------------------------------|------|
| POST     | `/api/register`         | Opret bruger                   | ❌   |
| POST     | `/api/login`            | Login og modtag JWT token      | ❌   |
| POST     | `/api/logout`           | Logout                         | ✅   |
| POST     | `/api/calculations`     | Evaluer matematisk udtryk      | ✅   |
| GET      | `/api/history`          | Hent seneste 5 beregninger     | ✅   |
| DELETE   | `/api/history/clear`    | Ryd historik                   | ✅   |