from locust import HttpUser, task, between

class BazarUser(HttpUser):
    wait_time = between(1, 5)  # Tiempo de espera entre tareas (1 a 5 segundos)

    @task(2)
    def list_products(self):
        self.client.get("/api/v1/products")

    @task(1)
    def create_product(self):
        self.client.post("/api/v1/products", json={
            "name": "Producto de prueba",
            "description": "Descripción de prueba",
            "category": "Categoría de prueba",
            "price": 10.99,
            "stock": 100
        })

    @task(1)
    def update_product(self):
        self.client.put("/api/v1/products/1", json={
            "name": "Producto actualizado",
            "description": "Descripción actualizada",
            "category": "Categoría actualizada",
            "price": 15.99,
            "stock": 50
        })

    @task(1)
    def delete_product(self):
        self.client.delete("/api/v1/products/1")

    def on_start(self):
        # Opcional: Código que se ejecuta al inicio de cada usuario
        pass