import random
from datetime import datetime
from locust import HttpUser, task, between
from faker import Faker

fake = Faker()


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
            "name": fake.word().capitalize() + " Producto",
            "description": fake.sentence(),
            "category": fake.word().capitalize(),
            "price": round(random.uniform(5.0, 100.0), 2),
            "stock": random.randint(1, 500)
        })
        if response.status_code == 201:
            product_id = response.json().get("product_id")
            if product_id:
                self.created_products.append(product_id)

    @task(1)
    def create_client(self):
        response = self.client.post("/api/v1/clients", json={
            "name": fake.name(),
            "email": fake.email(),
            "phone": fake.phone_number(),
            "address": fake.address(),
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
                "name": fake.word().capitalize() + " Actualizado",
                "description": fake.sentence(),
                "category": fake.word().capitalize(),
                "price": round(random.uniform(5.0, 100.0), 2),
                "stock": random.randint(1, 500)
            })

    @task(1)
    def update_client(self):
        if self.created_clients:
            client_id = random.choice(self.created_clients)
            self.client.put(f"/api/v1/clients/{client_id}", json={
                "name": fake.name(),
                "email": fake.email(),
                "phone": fake.phone_number(),
                "address": fake.address(),
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
            current_date = datetime.now().strftime("%Y-%m-%d")
            stock = random.randint(1, 5)
            self.client.post("/api/v1/sales", json={
                "client_id": client_id,
                "date": current_date,
                "products": [
                    {"product_id": product_id, "stock": stock}
                ],
                "quantity": stock,
                "total_amount": round(random.uniform(20.0, 200.0), 2)
            })

    def on_start(self):
        # Opcional: Crear un cliente y un producto inicial al inicio
        self.create_client()
        self.create_product()