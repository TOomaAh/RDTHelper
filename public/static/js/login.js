async function performLogin() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    const settings = {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    };

    try {
        const response = await fetch("/api/v1/login", settings);
        if (!response.ok) {
            throw new Error("Invalid username and password.");
        }
        const data = await response.json();
        const token = data.token;
        localStorage.setItem("token", token);
        window.location.href = "/web/home";
    } catch (error) {
        alert(error.message);
    }
}
