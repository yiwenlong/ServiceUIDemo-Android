package com.yiwenlong.service;

import android.os.Process;
import android.os.RemoteException;

import java.util.ArrayList;
import java.util.List;

import bridge.ILogHandler;
import bridge.MyServerListener;

public class RemoteServiceProxy extends IRemoteService.Stub {

    private List<MyServerListener> listeners = new ArrayList<>();

    private ILogHandler logHandler;

    @Override
    public void registerServiceCallback(IRemoteServiceCallback callback) {
        this.listeners.add(MyServerListenerImpl.newInstance(callback));
    }

    @Override
    public void registerLogCallback(IRemoteServiceLogCallback callback) {
        this.logHandler = LogHandlerImpl.newInstance(callback);
    }

    @Override
    public void bootServer(String host, int port) {
        bridge.Bridge.initServer(host, port);
        for (MyServerListener listener: listeners) {
            bridge.Bridge.registerServerListener(listener);
        }
        if (logHandler != null) {
            bridge.Bridge.registerLogHandler(logHandler);
        }
        bridge.Bridge.bootServer();
        listeners.clear();
        logHandler = null;
    }

    @Override
    public void stopServer() {
        bridge.Bridge.stopServer();
    }

    @Override
    public boolean isServerRunning() {
        return bridge.Bridge.serverIsRuning();
    }

    @Override
    public int getPid() {
        return Process.myPid();
    }

    static class LogHandlerImpl implements bridge.ILogHandler {

        private IRemoteServiceLogCallback remoteCallback;

        public static bridge.ILogHandler newInstance(IRemoteServiceLogCallback remoteCallback) {
            return new LogHandlerImpl(remoteCallback);
        }

        private LogHandlerImpl(IRemoteServiceLogCallback remoteCallback) {
            this.remoteCallback = remoteCallback;
        }

        @Override
        public void log(String s) {
            try {
                remoteCallback.onReceiveLog(s);
            } catch (RemoteException e) {
                e.printStackTrace();
            }
        }
    }

    static class MyServerListenerImpl implements bridge.MyServerListener {

        public static bridge.MyServerListener newInstance(IRemoteServiceCallback remoteCallback) {
            return new MyServerListenerImpl(remoteCallback);
        }

        private MyServerListenerImpl(IRemoteServiceCallback remoteCallback) {
            this.remoteCallback = remoteCallback;
        }

        private IRemoteServiceCallback remoteCallback;

        @Override
        public void onServerError(String msg) {
            try {
                remoteCallback.onError(msg);
            } catch (RemoteException e) {
                e.printStackTrace();
            }
        }

        @Override
        public void onServerStart() {
            try {
                remoteCallback.onStart();
            } catch (RemoteException e) {
                e.printStackTrace();
            }
        }

        @Override
        public void onServerStop() {
            try {
                remoteCallback.onStop();
            } catch (RemoteException e) {
                e.printStackTrace();
            }
        }
    }
}
