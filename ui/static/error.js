function deleteDefaultError() {
    setTimeout(function() {
        document.getElementById('signup-default-error').innerHTML = ''; 
        document.getElementById('login-default-error').innerHTML = ''; 
    }, 2000);
}