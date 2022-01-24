    //altert user notification
    function notifyUser(alertType) {
        switch (alertType) {
          case "success":
            notie.alert({
            type: "success", // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
            text: "form comleted successfully",
          })
            break;
            case "error":
            notie.alert({
            type: "error", // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
            text: "an error just occured",
          })
            break;
          default:
            notie.alert({
            type: "warning", // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
            text: "you not following the instructions well",
          })
            break;
        }
      }