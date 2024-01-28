function isEmailValid(email) {
    var emailRegex = /^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$/;
    return emailRegex.test(email);
}
function checkSignup() {
    var emailInput = document.getElementById('signupEmail');
    var errorMessageElement = document.getElementById('signup-email-message');

    if (!isEmailValid(emailInput.value)) {
        errorMessageElement.textContent = 'Неверная почта';
        setTimeout(function(){
            errorMessageElement.textContent = '';
        }, 1500);
    }
};

function checkLogin() {
    var emailInput = document.getElementById('loginEmail');
    var errorMessageElement = document.getElementById('login-email-message');

    if (!isEmailValid(emailInput.value)) {
        errorMessageElement.textContent = 'Неверная почта';
        setTimeout(function(){
            errorMessageElement.textContent = '';
        }, 1500);
    }
};

async function checkIfExists() {
    try {
        const response = await fetch('/auth/signup', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: document.getElementById('signupEmail').value,
                password: document.getElementById('signupPassword').value,
                name: document.getElementById('signupName').value
            })
        });

        const errorMsg = document.getElementById('signup-already');

        if (response.status === 500) {
            errorMsg.textContent = 'Пользователь с такой почтой уже существует';
            setTimeout(function() {
                errorMsg.textContent = '';
            }, 1500);
        }
    } catch (error) {
        console.error('Error:', error);
    }
};

async function checkPasswordAndEmail() {
    try {
        const response = await fetch('/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: document.getElementById('loginEmail').value,
                password: document.getElementById('loginPassword').value,
            })
        });

        const errorMsg = document.getElementById('invalid-pass-email');

        if (response.status === 500) {
            errorMsg.textContent = 'Неверная почта или пароль';
            setTimeout(function() {
                errorMsg.textContent = '';
            }, 1500);
        }
    } catch (error) {
        console.error('Error:', error);
    }
};