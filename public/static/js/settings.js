// before page load, check if token is present
window.addEventListener('load', function () {
    // check if token is present
    let token = localStorage.getItem('token');
    if (!token) {
        window.location.href = '/login';
    }

    // request /api/v1/settings to get settings
    let settings = {
        url: "/api/v1/settings",
        method: "GET",
    };

    makeRequest(settings, (response) => {
        // get username, password and read debrid token from html with id
        let username = document.getElementById("username");
        let password = document.getElementById("password");
        let debridToken = document.getElementById("rdt_token");

        // set the values of the inputs
        username.value = response.username;
        password.value = response.password;
        debridToken.value = response.rdt_token;
    });
});

function makeRequest(settings, callback) {
    // get token from local storage
    let token = localStorage.getItem('token');
    if (token) {
        settings.headers = {
            Authorization: 'Bearer ' + token
        };
    } else {
        // redirect to login page
        console.log("redirecting to login page");
        window.location.href = "/login";
        return;
    }

    fetch(settings.url, {
        method: settings.method,
        headers: settings.headers
    })
    .then(response => {
        if (response.headers.get('content-type').indexOf('text/html') >= 0) {
            window.location.href = "/login";
            return;
        }
        return response.json();
    })
    .then(data => callback(data))
    .catch(error => {
        console.log(error);
        let toastBody = document.getElementById('toast-body');
        if (toastBody.children.length === 0) {
            let json = JSON.parse(error.responseText);
            toastBody.innerHTML = `<p>${json.error}</p>`;
            let toast = new bootstrap.Toast(document.getElementById('liveToast'));
            toast.show();
        }
    });
}
