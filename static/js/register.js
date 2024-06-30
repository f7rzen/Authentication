const signupBtn = document.getElementById('signup-btn');

signupBtn.addEventListener('click', () => {
    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    fetch("/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ name, email, password })
    })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert(data.error);
            } else {
                // Перенаправление на страницу приветствия с именем пользователя
                window.location.href = `/welcome?name=${data.name}`;
            }
        })
        .catch(error => console.error("Error:", error));
});