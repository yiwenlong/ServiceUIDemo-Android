// IRemoteServiceCallback.aidl
package com.yiwenlong.service;

// Declare any non-default types here with import statements

interface IRemoteServiceCallback {

    void onStart();
    void onStop();
    void onError(String error);
}
