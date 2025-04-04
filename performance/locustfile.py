from locust import HttpUser, task, between
import random

class BazarUser(HttpUser):
    wait_time = between(1, 5)  # Tiempo de espera entre tareas (1 a 5 segundos)

    # Listas para almacenar IDs de clientes y productos creados
    created_clients = []
    created_products = []

    @task(2)
    def list_products(self):
        self.client.get("/api/v1/products")

    @task(2)
    def list_clients(self):
        self.client.get("/api/v1/clients")

    @task(1)
    def create_product(self):
        response = self.client.post("/api/v1/products", json={
            "name": "Producto de prueba",
            "description": "Descripción de prueba",
            "category": "Categoría de prueba",
            "price": 10.99,
            "stock": 100
        })
        if response.status_code == 201:
            product_id = response.json().get("product_id")
            if product_id:
                self.created_products.append(product_id)

    @task(1)
    def create_client(self):
        response = self.client.post("/api/v1/clients", json={
            "name": "Cliente de prueba",
            "email": "cliente@prueba.com",
            "phone": "123456789",
            "address": "Dirección de prueba"
        })
        if response.status_code == 201:
            client_id = response.json().get("client_id")
            if client_id:
                self.created_clients.append(client_id)

    @task(1)
    def update_product(self):
        if self.created_products:
            product_id = random.choice(self.created_products)
            self.client.put(f"/api/v1/products/{product_id}", json={
                "name": "Producto actualizado",
                "description": "Descripción actualizada",
                "category": "Categoría actualizada",
                "price": 15.99,
                "stock": 50
            })

    @task(1)
    def update_client(self):
        if self.created_clients:
            client_id = random.choice(self.created_clients)
            self.client.put(f"/api/v1/clients/{client_id}", json={
                "name": "Cliente actualizado",
                "email": "cliente@actualizado.com",
                "phone": "987654321",
                "address": "Dirección actualizada"
            })

    @task(1)
    def delete_product(self):
        if self.created_products:
            product_id = self.created_products.pop(0)  # Elimina el primer producto creado
            self.client.delete(f"/api/v1/products/{product_id}")

    @task(1)
    def delete_client(self):
        if self.created_clients:
            client_id = self.created_clients.pop(0)  # Elimina el primer cliente creado
            self.client.delete(f"/api/v1/clients/{client_id}")

    @task(1)
    def create_sale(self):
        if self.created_clients and self.created_products:
            client_id = random.choice(self.created_clients)
            product_id = random.choice(self.created_products)
            self.client.post("/api/v1/sales", json={
                "client_id": client_id,
                "date": "2025-04-04",
                "products": [
                    {"product_id": product_id, "stock": 2}
                ],
                "quantity": 2,
                "total_amount": 21.98
            })

    def on_start(self):
        # Opcional: Crear un cliente y un producto inicial al inicio
        self.create_client()
        self.create_product()