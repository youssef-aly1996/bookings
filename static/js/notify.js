    //altert user notification
    function notifyUser(alertType, msg) {
        switch (alertType) {
          case "success":
            notie.alert({
            type: "success", // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
            text: msg,
          })
            break;
            case "error":
            notie.alert({
            type: "error", // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
            text: msg,
          })
            break;
          default:
            notie.alert({
            type: "warning", // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
            text: msg,
          })
            break;
        }
      }