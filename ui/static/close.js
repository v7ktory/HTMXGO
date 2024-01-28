  // // Закрытие окна регистрации
  // document.addEventListener("DOMContentLoaded", function () {
  //   const closeBtn = document.querySelector("#signup #signup-close");
  //      closeBtn.addEventListener("click", () => {
  //      document.getElementById("signup").style.visibility = "hidden";
  //      document.getElementById("signupBtn").style.visibility = "visible";
  //  });

  //   const signupBtn = document.getElementById("signupBtn");
  //   signupBtn.addEventListener("click", () => {
  //         signupBtn.style.visibility = "hidden";
  //         document.getElementById("signup").style.visibility = "visible";
  //    });
  //  });
  //  document.getElementById("signup").style.visibility = "hidden";

  //    // Закрытие окна регистрации
  // document.addEventListener("DOMContentLoaded", function () {
  //   const closeBtn = document.querySelector("#login #login-close");
  //      closeBtn.addEventListener("click", () => {
  //      document.getElementById("login").style.visibility = "hidden";
  //      document.getElementById("loginBtn").style.visibility = "visible";
  //  });

  //   const signupBtn = document.getElementById("loginBtn");
  //   signupBtn.addEventListener("click", () => {
  //         signupBtn.style.visibility = "hidden";
  //         document.getElementById("login").style.visibility = "visible";
  //    });
  //  });
  //  document.getElementById("login").style.visibility = "hidden";

//   document.addEventListener("DOMContentLoaded", function () {
//     const signupBtn = document.getElementById("signupBtn");
//     const signupCloseBtn = document.querySelector("#signup #signup-close");
//     const loginBtn = document.getElementById("loginBtn");
//     const loginCloseBtn = document.querySelector("#login #login-close");

//     signupBtn.addEventListener("click", () => {
//         toggleVisibility("signup", "signupBtn");
//         closeIfVisible("login", "loginBtn");
//     });

//     signupCloseBtn.addEventListener("click", () => {
//         toggleVisibility("signup", "signupBtn");
//     });

//     loginBtn.addEventListener("click", () => {
//         toggleVisibility("login", "loginBtn");
//         closeIfVisible("signup", "signupBtn");
//     });

//     loginCloseBtn.addEventListener("click", () => {
//         toggleVisibility("login", "loginBtn");
//     });
// });

// function toggleVisibility(elementId, btnId) {
//     const element = document.getElementById(elementId);
//     const btn = document.getElementById(btnId);

//     if (element.style.visibility === "visible") {
//         element.style.visibility = "hidden";
//         btn.style.visibility = "visible";
//     } else {
//         element.style.visibility = "visible";
//         btn.style.visibility = "hidden";
//     }
// }

// function closeIfVisible(elementId, btnId) {
//     const element = document.getElementById(elementId);
//     const btn = document.getElementById(btnId);

//     if (element.style.visibility === "visible") {
//         element.style.visibility = "hidden";
//         btn.style.visibility = "visible";
//     }
// }

// // Скрытие окон при загрузке страницы
// document.getElementById("signup").style.visibility = "hidden";
// document.getElementById("login").style.visibility = "hidden";

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
      element.style.zIndex = "5"; // Устанавливаем более высокий z-index при отображении
  }
}

function closeIfVisible(elementId, btnId) {
  const element = document.getElementById(elementId);
  const btn = document.getElementById(btnId);

  if (element.style.visibility === "visible") {
      element.style.visibility = "hidden";
      btn.style.visibility = "visible";
      element.style.zIndex = "0"; // Задаем низкий z-index при скрытии
  }
}

// Скрытие окон при загрузке страницы
document.getElementById("signup").style.visibility = "hidden";
document.getElementById("login").style.visibility = "hidden";

