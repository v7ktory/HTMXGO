// Проверка почты
function isEmailValid(email) {
    var emailRegex = /^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$/;
    return emailRegex.test(email);
}
function checkSignup() {
    var emailInput = document.getElementById('signupEmail');
    var errorMessageElement = document.getElementById('signup-error-message');

    if (!isEmailValid(emailInput.value)) {
        errorMessageElement.textContent = 'Неверная почта';
    }
};

function checkLogin() {
    var emailInput = document.getElementById('loginEmail');
    var errorMessageElement = document.getElementById('login-error-message');

    if (!isEmailValid(emailInput.value)) {
        errorMessageElement.textContent = 'Неверная почта';
    }
};