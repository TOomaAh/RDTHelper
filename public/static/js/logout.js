function logout() {
    // Remove the token from the local storage
    localStorage.removeItem('token');
    // Redirect to the login page
    window.location.href = "/login";
    // Return false to prevent the default behavior of the event
    return false;
}