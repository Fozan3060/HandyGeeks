let isSideMenuVisible = false;
function toggleSideMenuVisbility() {
  if (isSideMenuVisible) {
    document.querySelector("aside").style.display = "none";
    document.querySelector(".overlay").style.display = "none";
  } else {
    document.querySelector("aside").style.display = "block";
    document.querySelector(".overlay").style.display = "block";
  }
  isSideMenuVisible = !isSideMenuVisible;
}

document.addEventListener("DOMContentLoaded", function () {
  // contact us form
  const loginForm = document.getElementById("login-form");

  loginForm.addEventListener("submit", async function (event) {
    event.preventDefault(); // Prevent the default form submission behavior

    const formData = new FormData(loginForm);
    const contactUsData = {};

    formData.forEach((value, key) => {
      contactUsData[key] = value;
    });

    console.log("User data:", contactUsData);

    try {
      const response = await fetch("http://localhost:3000/contact-us", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(contactUsData),
      });

      console.log("Status Code Received :", response.status);
      if (response.status === 200) {
        alert("Thank You for contacting us. We will get back to you soon.");
      } else {
        alert("Something went wrong. Please try again later.");
      }
    } catch (error) {
      console.log(error);
    }
  });
});
