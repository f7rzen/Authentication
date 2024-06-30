const signupBtn = document.querySelector('.search-box button');

signupBtn.addEventListener('click', () => {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    fetch("/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ email, password })
    })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert(data.error);
            } else {
                // Перенаправление на страницу приветствия
                window.location.href = `/welcome?name=${data.name}`;
            }
        })
        .catch(error => console.error("Error:", error));
});