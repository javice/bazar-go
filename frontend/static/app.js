document.addEventListener('DOMContentLoaded', () => {
    const apiUrl = 'https://bazar-go.onrender.com/api/v1'; 
    const productosDiv = document.getElementById('productos');
    const formProducto = document.getElementById('formProducto');
    const formVenta = document.getElementById('formVenta');
    const addProductRowButton = document.getElementById('addProductRow');
    const productRowsContainer = document.getElementById('productRows');

    // Cargar productos al iniciar
    fetch(`${apiUrl}/products`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Error en la API: ${response.statusText}`);
            }
            return response.json();
        })
        .then(data => {
            productosDiv.innerHTML = data.data.map(product => `
                <div class="producto" data-id="${product.product_id}">
                    <h3>${product.name}</h3>
                    <p>Descripción: ${product.description || "Sin descripción"}</p>
                    <p>Categoría: ${product.category || "Sin categoría"}</p>
                    <p>Precio: $${product.price.toFixed(2)}</p>
                    <p>Stock: ${product.stock}</p>
                    <div class="producto-actions">
                        <button onclick='editProduct(${JSON.stringify(product)})' class="btn-edit">Editar</button>
                        <button onclick="deleteProduct(${product.product_id})" class="btn-delete">Eliminar</button>
                    </div>
                </div>
            `).join('');
        })
        .catch(error => {
            console.error("Error al cargar los productos:", error);
            productosDiv.innerHTML = `<p>Error al cargar los productos.</p>`;
        });

    // Agregar nuevo producto
    formProducto.addEventListener('submit', (e) => {
        e.preventDefault();
        const nombre = document.getElementById('nombre').value;
        const precio = parseFloat(document.getElementById('precio').value);
        const categoria = document.getElementById('categoria').value;
        const descripcion = document.getElementById('descripcion').value;
        const stock = parseInt(document.getElementById('stock').value, 10);

        fetch(`${apiUrl}/products`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name: nombre, price: precio, category: categoria, description: descripcion, stock: stock })
        })
        .then(() => location.reload()); // Recargar para ver cambios
    });

    // Eliminar producto
    window.deleteProduct = async (id) => {
        if (confirm("¿Eliminar este producto?")) {
            await fetch(`${apiUrl}/products/${id}`, { method: 'DELETE' });
            location.reload();
        }
    };

    // Función para abrir modal de edición
    window.editProduct = (product) => {
        if (typeof product === 'string') {
            product = JSON.parse(product); // Convertir de cadena JSON a objeto
        }
        document.getElementById('editId').value = product.product_id;
        document.getElementById('editName').value = product.name;
        document.getElementById('editDescription').value = product.description;
        document.getElementById('editCategory').value = product.category;
        document.getElementById('editStock').value = product.stock;
        document.getElementById('editPrice').value = product.price;
        document.getElementById('editModal').style.display = 'block';
    };

    // Cerrar modal al hacer clic en la "X"
    document.querySelector('.close').addEventListener('click', () => {
        document.getElementById('editModal').style.display = 'none';
    });

    // Función para enviar la actualización
    document.getElementById('editForm').addEventListener('submit', async (e) => {
        e.preventDefault();
        const id = document.getElementById('editId').value;
        const name = document.getElementById('editName').value;
        const description = document.getElementById('editDescription').value;
        const category = document.getElementById('editCategory').value;
        const stock = document.getElementById('editStock').value;
        const price = parseFloat(document.getElementById('editPrice').value);

        await fetch(`${apiUrl}/products/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ category,description,name, price,stock })
        });
        location.reload();
    });

    // Agregar nueva fila de producto en el formulario de ventas
    addProductRowButton.addEventListener('click', () => {
        const newRow = document.createElement('div');
        newRow.classList.add('productRow');

        newRow.innerHTML = `
            <label for="product_id">ID del Producto:</label>
            <input type="number" name="product_id[]" placeholder="ID Producto" required>

            <label for="cantidad">Cantidad:</label>
            <input type="number" name="cantidad[]" placeholder="Cantidad" step="1" required>
        `;

        productRowsContainer.appendChild(newRow);
    });

    // Abrir el modal para confirmar la venta
    formVenta.addEventListener('submit', (e) => {
        e.preventDefault(); // Prevenir el comportamiento predeterminado del formulario

        const clientId = parseInt(document.getElementById('client_id').value, 10);

        // Obtener la fecha actual en formato YYYY-MM-DD
        const date = new Date().toISOString().split('T')[0];

        // Obtener los productos y cantidades
        const productIds = Array.from(document.getElementsByName('product_id[]')).map(input => parseInt(input.value, 10));
        const quantities = Array.from(document.getElementsByName('cantidad[]')).map(input => parseInt(input.value, 10));

        // Validar que no haya valores NaN o negativos
        if (productIds.some(isNaN) || quantities.some(isNaN) || quantities.some(q => q <= 0)) {
            alert("Por favor, ingresa valores válidos para los productos y cantidades.");
            return;
        }

        // Calcular la suma total de las cantidades
        const totalQuantity = quantities.reduce((sum, quantity) => sum + quantity, 0);

        // Crear el array de productos para la venta
        const products = productIds.map((id, index) => ({
            product_id: id,
            stock: quantities[index]
        }));

        // Mostrar los datos en el modal
        const ventaResumen = document.getElementById('ventaResumen');
        ventaResumen.innerHTML = `
            <p><strong>ID Cliente:</strong> ${clientId}</p>
            <p><strong>Fecha:</strong> ${date}</p>
            <p><strong>Productos:</strong></p>
            <ul>
                ${products.map(product => `<li>ID: ${product.product_id}, Cantidad: ${product.stock}</li>`).join('')}
            </ul>
            <p><strong>Total de Productos:</strong> ${totalQuantity}</p>
        `;

        // Guardar los datos en variables globales para usarlos al confirmar la venta
        window.ventaData = { client_id: clientId, date, products, quantity: totalQuantity };

        // Abrir el modal
        document.getElementById('ventaModal').style.display = 'block';
    });

    // Cerrar el modal al hacer clic en la "X"
    document.getElementById('closeVentaModal').addEventListener('click', () => {
        document.getElementById('ventaModal').style.display = 'none';
    });

    // Confirmar la venta al hacer clic en "Confirmar Venta"
    document.getElementById('confirmarVenta').addEventListener('click', () => {
        const { client_id, date, products, quantity } = window.ventaData;

        // Enviar la venta al backend
        fetch(`${apiUrl}/sales`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ client_id, date, products, quantity })
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Error al generar la venta: ${response.statusText}`);
            }
            return response.json();
        })
        .then(() => {
            alert("Venta generada correctamente");
            document.getElementById('ventaModal').style.display = 'none';
            location.reload();
        })
        .catch(error => {
            console.error("Error al generar la venta:", error);
            alert("Hubo un error al generar la venta.");
        });
    });
});


