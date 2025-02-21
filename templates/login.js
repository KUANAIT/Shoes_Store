document.addEventListener('DOMContentLoaded', function () {
    const loginButton = document.getElementById('loginButton');
    const messageDiv = document.getElementById('message');

    loginButton.addEventListener('click', function (event) {
        event.preventDefault();

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        fetch('http://localhost:3000/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        })
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else {
                    throw new Error('Login failed');
                }
            })
            .then(data => {
                messageDiv.style.display = 'block';
                messageDiv.innerHTML = `<h2>Authentication Success</h2><p>Welcome back, ${username}</p>`;

                setTimeout(() => {
                    window.location.href = '/';
                }, 2000);
            })
            .catch(error => {
                messageDiv.style.display = 'block';
                messageDiv.innerHTML = `<h2>Error</h2><p>${error.message}</p>`;
            });
    });
});