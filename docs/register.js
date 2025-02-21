document.addEventListener('DOMContentLoaded', function () {
    const registerButton = document.getElementById('registerButton');
    const messageDiv = document.getElementById('message');
    const errorDiv = document.getElementById('error');

    registerButton.addEventListener('click', function (event) {
        event.preventDefault();

        const username = document.getElementById('username').value;
        const email = document.getElementById('email').value;
        const firstName = document.getElementById('firstName').value;
        const lastName = document.getElementById('lastName').value;
        const shippingAddress = document.getElementById('shippingAddress').value;
        const phone = document.getElementById('phone').value;
        const password = document.getElementById('password').value;

        fetch('http://localhost:3000/user/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                email: email,
                first_name: firstName,
                last_name: lastName,
                shipping_address: shippingAddress,
                phone: phone,
                password: password,
            }),
        })
            .then(response => {
                console.log("Response status:", response.status);
                if (response.ok) {
                    return response.json().then(data => {
                        console.log("Response data:", data);
                        return data;
                    });
                } else {
                    return response.json().then(err => {
                        console.log("Error response:", err);
                        throw new Error(err.message || 'Registration failed');
                    });
                }
            })
            .then(data => {
                window.location.href = '/loginuser';
            })
            .catch(error => {
                console.error("Fetch error:", error);
                errorDiv.style.display = 'block';
                errorDiv.innerHTML = `<h2>Error</h2><p>${error.message}</p>`;
            });
    });
});