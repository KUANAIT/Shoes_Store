document.addEventListener('DOMContentLoaded', function() {
  fetchShoes();
});

function fetchShoes() {
  fetch('http://localhost:3000/getall')
      .then(response => response.json())
      .then(data => {
        displayShoes(data);
      })
      .catch(error => {
        console.error('Error fetching shoes:', error);
      });
}

function displayShoes(shoes) {
  const shoeList = document.getElementById('shoeList');
  shoeList.innerHTML = '';

  shoes.forEach(shoe => {
    const shoeItem = document.createElement('div');
    shoeItem.className = 'galleryItem';

    const shoeId = document.createElement('p');
    shoeId.textContent = `ID: ${shoe.id}`;

    const shoeName = document.createElement('h2');
    shoeName.textContent = shoe.name;

    const shoeBrand = document.createElement('p');
    shoeBrand.textContent = `Brand: ${shoe.brand}`;

    const shoePrice = document.createElement('p');
    shoePrice.textContent = `Price: ₸${shoe.price}`;

    const shoeDiscount = document.createElement('p');
    shoeDiscount.textContent = `Discount: ${shoe.discount}%`;

    const shoeSize = document.createElement('p');
    shoeSize.textContent = `Sizes: ${shoe.size.join(', ')}`;

    const shoeMaterial = document.createElement('p');
    shoeMaterial.textContent = `Material: ${shoe.material}`;

    const shoeColor = document.createElement('p');
    shoeColor.textContent = `Colors: ${shoe.color.join(', ')}`;

    const shoeReleaseDate = document.createElement('p');
    shoeReleaseDate.textContent = `Release Date: ${shoe.release_date}`;

    const shoeStockQuantity = document.createElement('p');
    shoeStockQuantity.textContent = `Stock Quantity: ${shoe.stock_quantity}`;

    const shoeRating = document.createElement('p');
    shoeRating.textContent = `Rating: ${shoe.rating}/5`;

    const shoeFeatures = document.createElement('p');
    shoeFeatures.textContent = `Features: ${shoe.features.join(', ')}`;

    const shoeCategory = document.createElement('p');
    shoeCategory.textContent = `Category: ${shoe.category}`;

    const shoeShippingInfo = document.createElement('p');
    shoeShippingInfo.textContent = `Shipping: ${shoe.shipping_info.free_shipping ? 'Free Shipping' : 'Paid Shipping'}`;

    shoeItem.appendChild(shoeId);
    shoeItem.appendChild(shoeName);
    shoeItem.appendChild(shoeBrand);
    shoeItem.appendChild(shoePrice);
    shoeItem.appendChild(shoeDiscount);
    shoeItem.appendChild(shoeSize);
    shoeItem.appendChild(shoeMaterial);
    shoeItem.appendChild(shoeColor);
    shoeItem.appendChild(shoeReleaseDate);
    shoeItem.appendChild(shoeStockQuantity);
    shoeItem.appendChild(shoeRating);
    shoeItem.appendChild(shoeFeatures);
    shoeItem.appendChild(shoeCategory);
    shoeItem.appendChild(shoeShippingInfo);

    shoeList.appendChild(shoeItem);
  });
}

document.addEventListener('DOMContentLoaded', function() {
  fetchShoes();

  const orderForm = document.getElementById('orderForm');
  orderForm.addEventListener('submit', function(event) {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const shoeId = document.getElementById('shoeId').value;
    const quantity = parseInt(document.getElementById('quantity').value);

    fetch(`http://localhost:3000/getbyid?id=${shoeId}`)
        .then(response => response.json())
        .then(shoe => {
          const order = {
            username: username,
            order: {
              user_id: "",
              order_date: new Date().toISOString().split('T')[0],
              shipping_address: "123 Main St",
              items: [{
                shoe_id: shoeId,
                quantity: quantity,
                price: shoe.price
              }],
              total_price: shoe.price * quantity,
              status: "Pending",
              shipping_info: {
                free_shipping: false,
                expedited_shipping: "Standard"
              }
            }
          };

          fetch('http://localhost:3000/order/createforuser', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(order)
          })
              .then(response => response.json())
              .then(data => {
                alert(`Order created successfully! Order ID: ${data.id}`);
              })
              .catch(error => {
                console.error('Error creating order:', error);
              });
        })
        .catch(error => {
          console.error('Error fetching shoe details:', error);
        });
  });
});
