document.addEventListener("DOMContentLoaded", function () {
  const signupBtn = document.getElementById("signupBtn");
  const signupCloseBtn = document.querySelector("#signup #signup-close");
  const loginBtn = document.getElementById("loginBtn");
  const loginCloseBtn = document.querySelector("#login #login-close");

  signupBtn.addEventListener("click", () => {
      toggleVisibility("signup", "signupBtn");
      closeIfVisible("login", "loginBtn");
  });

  signupCloseBtn.addEventListener("click", () => {
      toggleVisibility("signup", "signupBtn");
  });

  loginBtn.addEventListener("click", () => {
      toggleVisibility("login", "loginBtn");
      closeIfVisible("signup", "signupBtn");
  });

  loginCloseBtn.addEventListener("click", () => {
      toggleVisibility("login", "loginBtn");
  });
});

function toggleVisibility(elementId, btnId) {
  const element = document.getElementById(elementId);
  const btn = document.getElementById(btnId);

  if (element.style.visibility === "visible") {
      element.style.visibility = "hidden";
      btn.style.visibility = "visible";
  } else {
      element.style.visibility = "visible";
      btn.style.visibility = "hidden";
      element.style.zIndex = "5"; 
  }
}

function closeIfVisible(elementId, btnId) {
  const element = document.getElementById(elementId);
  const btn = document.getElementById(btnId);

  if (element.style.visibility === "visible") {
      element.style.visibility = "hidden";
      btn.style.visibility = "visible";
      element.style.zIndex = "0"; 
  }
}

document.getElementById("signup").style.visibility = "hidden";
document.getElementById("login").style.visibility = "hidden";

