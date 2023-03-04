// static/js/script.js
function validateTimetableForm() {
  // Get form elements by id
  var subject = document.getElementById("subject");
  var startDay = document.getElementById("startDay");
  var endDay = document.getElementById("endDay");

  // Check if subject is empty
  if (subject.value == "") {
    alert("Please enter a subject.");
    return false;
  }

  // Check if start day is valid
  if (isNaN(startDay.value) || startDay.value < 1 || startDay.value > 31) {
    alert("Please enter a valid start day (1-31).");
    return false;
  }

  // Check if end day is valid
  if (isNaN(endDay.value) || endDay.value < startDay.value || endDay.value > 31) {
    alert("Please enter a valid end day (1-31) that is greater than or equal to start day.");
    return false;
  }

  // If no error, return true
  return true;
}
