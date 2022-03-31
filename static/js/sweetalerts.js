let success = function (c) {
  const { msg = "", title = "", footer = "", showConfirmButton = true } = c;

  Swal.fire({
    icon: "success",
    title: title,
    html: msg,
    showConfirmButton: showConfirmButton,
    footer: footer,
  });
};

let error = function (c) {
  const { msg = "", title = "", footer = "" } = c;

  Swal.fire({
    icon: "error",
    title: title,
    text: msg,
    footer: footer,
  });
};
async function custom(c) {
  const { msg = "", title = "", icon = "", showConfirmButton = true } = c;
  const { value: formValues } = await Swal.fire({
    icon: icon,
    title: title,
    html: msg,
    focusConfirm: false,
    showCancelButton: true,
    showConfirmButton: showConfirmButton,
    willOpen: () => {
      const elem = document.getElementById("reservation-dates-modal");
      const rangePicker = new DateRangePicker(elem, {
        format: "yyyy-mm-dd",
        minDate: new Date(),
      });
    },
    preConfirm: () => {
      return [
        document.getElementById("start").value,
        document.getElementById("end").value,
      ];
    },
    didOpen: () => {
      document.getElementById("start").removeAttribute("disabled");
      document.getElementById("end").removeAttribute("disabled");
    },
  });

  if (formValues) {
    if (formValues.dismiss !== Swal.DismissReason.cancel) {
      if (formValues !== "") {
        if (c.callback !== undefined) {
          c.callback(formValues);
        } else {
          c.callback(false);
        }
      }
    } else {
      c.callback(false);
    }
  }
}

document
  .getElementById("search-availability-button")
  .addEventListener("click", () => {
    let html = `
      <div class="col">
          <form action="" method="post" id="search-availability-form" novalidate class="needs-validation">
            <div class="form-row">
              <div class="col">
                <div class="form-row" id="reservation-dates-modal">
                  <div class="col">
                    <input disabled type="text" name="start" id="start" class="form-control" placeholder="Arrival" autocomplete="off">
                  </div>
                  <div class="col">
                    <input disabled type="text" name="end" id="end" class="form-control" placeholder="Departure" autocomplete="off">
                  </div>
                </div>
              </div>
            </div>
          </form>
      </div>
        `;
    custom({
      msg: html,
      title: "choose the date",
      callback: async function (formValues) {
        let form = document.getElementById("search-availability-form");
        let formData = new FormData(form);
        formData.append("csrf_token", "{{.CSRF}}");
        formData.append("room_id", "1");
        const response = await fetch("/search-availability-json", {
          method: "post",
          body: formData,
        });
        let json = await response.json();
        if (json.ok) {
          success({
            icon: "success",
            showConfirmButton: false,
            msg: `<p> room is available </p>
              <p> <a href="/book-room?id=${json.room_id}&sd=${json.start_date}&ed=${json.end_date}" class="btn btn-primary">
              Book Now! </a></p`,
          });
        } else {
          error({
            msg: "not available",
          });
        }
      },
    });
  });
