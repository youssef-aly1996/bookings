      async function custom(c) {
        const{
          msg = "",
          title = ""
        } = c;
        const { value: formValues } = await Swal.fire({
                title: title,
                html:msg,
                focusConfirm: false,
                showCancelButton: true,
                willOpen: () => {
                  const elem = document.getElementById("search-availability-form");
                  const rangePicker = new DateRangePicker(elem, {
                    format: "yyyy-mm-dd",
                  });
                },
                preConfirm: () => {
                  return [
                  document.getElementById('start').value,
                  document.getElementById('end').value
                  ]
                },
                didOpen: () => {
                  document.getElementById('start').removeAttribute('disabled');
                  document.getElementById('end').removeAttribute('disabled');
                },
              });

              if (formValues) {
                Swal.fire(JSON.stringify(formValues))
            }
}


document.getElementById("search-availability-button").addEventListener("click", () => {
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
  custom({msg:html, title: "choose the date"});
  });