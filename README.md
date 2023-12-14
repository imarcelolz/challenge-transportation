# Transportation Management Challenge

## Problem Description

**Design/implement a system to manage transportation for multiple groups of people traveling to the same destination.**

Given a set of cars with varying seat capacities and a set of groups of people requesting transportation, the system should efficiently manage the allocation of groups to cars based on their size. The system should prioritize serving groups as quickly as possible while maintaining the order of arrival when possible. It should also be able to handle a large number of cars and groups while ensuring high performance and scalability.

## System Requirements

1. **Efficient Seat Management:** The system should effectively manage the available seats in cars and ensure that groups are assigned to cars based on their size without overfilling any car.

2. **Fairness and Priority:** Groups should be prioritized for assignment to cars based on their arrival time, ensuring that groups that arrive earlier are served before those that arrive later, whenever possible.

3. **Scalability:** The system should be able to handle a large number of cars and groups, ideally reaching at least $10^4$ cars and $10^5$ groups. This requires efficient data structures and algorithms to optimize resource utilization.

4. **Performance:** The system should maintain high performance even when handling a large number of cars and groups. This entails using optimized data structures and algorithms, along with efficient implementation techniques.

5. **Testing and Robustness:** The system should be extensively tested to ensure its correctness and robustness. It should handle a variety of error conditions gracefully and maintain stability under load.

6. **REST API:** The system should provide a REST API for managing the transportation of groups of people. The API should include the following endpoints:

    a. **GET /status:** This endpoint should return a response indicating the status of the system, such as the number of available cars, the number of waiting groups, and the overall system health.

    b. **PUT /cars:** This endpoint should load the list of available cars in the system and reset the application state. The request body should be a JSON array of objects representing the cars, each with the following properties:
        1. `id`: The unique identifier of the car.
        2. `seats`: The number of seats available in the car.

    c. **POST /journey:** This endpoint should register a group of people requesting transportation. The request body should be a JSON object representing the group, with the following properties:
        1. `id`: The unique identifier of the group.
        2. `people`: The number of people in the group.

    d. **POST /dropoff:** This endpoint should indicate that a group of people has completed their journey. The request body should be a URL-encoded form with the group ID, such that `ID=X`.

    e. **POST /locate:** This endpoint should return the car that a group of people is traveling with, or no car if they are still waiting to be assigned. The request body should be a URL-encoded form with the group ID, such that `ID=X`. The response should be a JSON object representing the car, or an empty object if the group is not assigned to a car.
