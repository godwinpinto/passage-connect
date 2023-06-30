#include "string.h"
#include "stdlib.h"
#include "stdio.h"
#include "_cgo_export.h"

#define PAM_SM_AUTH
#define PAM_SM_PASSWORD
#include <security/pam_appl.h>
#include <security/pam_modules.h>
#include <security/pam_ext.h>

GoSlice argcvToSlice(int, const char**);

PAM_EXTERN int pam_sm_open_session(pam_handle_t *pamh, int flags, int argc, const char **argv) {
    // Perform any necessary operations when opening a session
    pam_info(pamh, "Authentication in progress 1...");

    const struct pam_conv *conv;
    struct pam_message msg;
    const struct pam_message *msgp;
    struct pam_response *resp;

    int retval = pam_get_item(pamh, PAM_CONV, (const void **)&conv);
    if (retval != PAM_SUCCESS || conv == NULL) {
        return PAM_SYSTEM_ERR;
    }

    msg.msg_style = PAM_PROMPT_ECHO_OFF;
    msg.msg = "Enter your password: ";
    msgp = &msg;

    retval = conv->conv(1, &msgp, &resp, conv->appdata_ptr);
    if (retval != PAM_SUCCESS) {
        return retval;
    }
    
    printf("Passage Session opened successfully!\n");
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


