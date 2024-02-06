document.addEventListener("DOMContentLoaded", function () {
    const profileBtn = document.getElementById("profileBtn");
    const profileCloseBtn = document.querySelector("#profile-info #profile-close");

    profileBtn.addEventListener("click", () => {
        toggleVisibility("profile-info", "profileBtn");
    });
  
    profileCloseBtn.addEventListener("click", () => {
        toggleVisibility("profile-info", "profileBtn");
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
  
document.getElementById("profile-info").style.visibility = "hidden";