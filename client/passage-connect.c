#include "string.h"
#include "stdlib.h"
#include "stdio.h"
#include "_cgo_export.h"

#define PAM_SM_AUTH
#define PAM_SM_PASSWORD
#include <security/pam_appl.h>
#include <security/pam_modules.h>

GoSlice argcvToSlice(int, const char**);

static int my_conv(int num_msg, const struct pam_message **msg, struct pam_response **resp, void *appdata_ptr) {
    if (num_msg <= 0 || msg == NULL || resp == NULL) {
        return PAM_CONV_ERR;
    }

    // Allocate memory for the response array
    struct pam_response *response = (struct pam_response *)malloc(sizeof(struct pam_response) * num_msg);
    if (response == NULL) {
        return PAM_CONV_ERR;
    }

    // Process each message
    for (int i = 0; i < num_msg; i++) {
        const struct pam_message *message = msg[i];
        struct pam_response *r = &response[i];

        if (message->msg_style == PAM_PROMPT_ECHO_OFF) {
            // Prompt for password without echoing
            r->resp = strdup("mypassword");
        } else if (message->msg_style == PAM_TEXT_INFO) {
            // Display informational message
            printf("%s\n", message->msg);
            r->resp = NULL;
        } else {
            // Unsupported message style
            free(response);
            return PAM_CONV_ERR;
        }
    }

    *resp = response;
    return PAM_SUCCESS;
}

PAM_EXTERN int pam_sm_open_session(pam_handle_t *pamh, int flags, int argc, const char **argv) {
    // Perform any necessary operations when opening a session
    printf("Passage Session opened successfully!\n");
    pam_set_item(pamh, PAM_CONV, (const void *)&my_conv);
    return goAuthenticate(pamh, flags, argcvToSlice(argc, argv));
    //return PAM_SUCCESS;
}

PAM_EXTERN int pam_sm_close_session(pam_handle_t *pamh, int flags, int argc, const char **argv) {
    // Perform any necessary operations when closing a session

    printf("Passage Session closed successfully!\n");
    return PAM_SUCCESS;
}

PAM_EXTERN int pam_sm_authenticate(pam_handle_t* pamh, int flags, int argc, const char** argv) {
    printf("Passage Connect Auth called!\n");
  return goAuthenticate(pamh, flags, argcvToSlice(argc, argv));
}

PAM_EXTERN int pam_sm_setcred(pam_handle_t* pamh, int flags, int argc, const char** argv) {
    printf("Set Cred called!\n");
  return setCred(pamh, flags, argcvToSlice(argc, argv));
}

GoSlice argcvToSlice(int argc, const char** argv) {
  GoString* strs = malloc(sizeof(GoString) * argc);

  GoSlice ret;
  ret.cap = argc;
  ret.len = argc;
  ret.data = (void*)strs;

  int i;
  for(i = 0; i < argc; i++) {
    strs[i] = *((GoString*)malloc(sizeof(GoString)));

    strs[i].p = (char*)argv[i];
    strs[i].n = strlen(argv[i]);
  }

  return ret;
}

void myCFunction(const char* str) {
    printf("Hello from C: %s\n", str);
}

