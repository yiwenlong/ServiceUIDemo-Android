package com.yiwenlong.serviceui;

import androidx.appcompat.app.AppCompatActivity;

import android.content.ComponentName;
import android.content.Context;
import android.content.Intent;
import android.content.ServiceConnection;
import android.os.Bundle;
import android.os.IBinder;
import android.os.RemoteException;
import android.util.Log;
import android.view.View;
import android.widget.Button;

import com.yiwenlong.service.IRemoteService;
import com.yiwenlong.service.IRemoteServiceCallback;
import com.yiwenlong.service.IRemoteServiceLogCallback;
import com.yiwenlong.service.ServerManagerService;

public class MainActivity extends AppCompatActivity {
    private static final String TAG = MainActivity.class.getSimpleName();

    private IRemoteService iRemoteService;

    private Button bind, unbind, start, stop;

    private ServiceConnection mConnection = new ServiceConnection() {

        // Called when the connection with the service is established
        public void onServiceConnected(ComponentName className, IBinder service) {
            // Following the example above for an AIDL interface,
            // this gets an instance of the IRemoteInterface, which we can use to call on the service
            iRemoteService = IRemoteService.Stub.asInterface(service);

            try {
                iRemoteService.registerLogCallback(logCallback);
                iRemoteService.registerServiceCallback(serviceCallback);
            } catch (RemoteException e) {
                e.printStackTrace();
            }

            // enable start button.
            if (bind != null) {
                bind.setEnabled(false);
            }

            if (start != null) {
                start.setEnabled(true);
            }
            if (unbind != null) {
                unbind.setEnabled(true);
            }
        }

        // Called when the connection with the service disconnects unexpectedly
        public void onServiceDisconnected(ComponentName className) {
            Log.e(TAG, "Service has unexpectedly disconnected");
            iRemoteService = null;
        }
    };

    private IRemoteServiceLogCallback logCallback = new IRemoteServiceLogCallback.Stub() {
        @Override
        public void onReceiveLog(String message) {
            Log.i(TAG, message);
        }
    };

    private IRemoteServiceCallback serviceCallback = new IRemoteServiceCallback.Stub() {
        @Override
        public void onStart() {
            Log.i(TAG, "service -> onStart");
        }

        @Override
        public void onStop() {
            Log.i(TAG, "service -> onStop");
        }

        @Override
        public void onError(String error) {
            Log.i(TAG, "service -> onError: " + error);
        }
    };

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        bind = findViewById(R.id.bind);
        unbind = findViewById(R.id.unbind);
        start = findViewById(R.id.start);
        stop = findViewById(R.id.stop);

        bind.setOnClickListener(view -> {
            Intent i = new Intent(getApplicationContext(), ServerManagerService.class);
            bindService(i, mConnection, Context.BIND_AUTO_CREATE);
        });

        unbind.setOnClickListener(view -> {
            unbindService(mConnection);
            bind.setEnabled(true);
            start.setEnabled(false);
            stop.setEnabled(false);
        });

        start.setOnClickListener(view -> {
            try {
                iRemoteService.bootServer("0.0.0.0", 8080);
            } catch (RemoteException e) {
                e.printStackTrace();
            }
            stop.setEnabled(true);
            start.setEnabled(false);
        });

        stop.setOnClickListener(view -> {
            try {
                iRemoteService.stopServer();
            } catch (RemoteException e) {
                e.printStackTrace();
            }
            stop.setEnabled(false);
            start.setEnabled(true);
        });
    }
}