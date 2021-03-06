//
// Copyright 2020 Yiwenlong(wlong.yi#gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may obtain a copy of the License at
// you may not use this file except in compliance with the License.
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package com.yiwenlong.service;

import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.Service;
import android.content.Intent;
import android.os.IBinder;
import android.util.Log;

import androidx.annotation.Nullable;
import androidx.core.app.NotificationCompat;
import androidx.core.app.NotificationManagerCompat;

public class ServerManagerService extends Service {

    private static final String CHANNEL_ID = "channel id";

    @Override
    public void onCreate() {
        Log.i("ServerManagerService", "onCreate");
        super.onCreate();

        NotificationManagerCompat.from(this)
                .createNotificationChannel(
                        new NotificationChannel(
                                CHANNEL_ID,
                                "a demo channel",
                                NotificationManager.IMPORTANCE_DEFAULT
                        )
                );

        Notification notification = new NotificationCompat.Builder(getApplicationContext(),CHANNEL_ID)
                .setContentText("这是一个 Android 下运行 Golang server 的 App")
                .setContentTitle("ServiceUI Shell")
                .setSmallIcon(R.mipmap.ic_launcher_round)
                .setGroup(CHANNEL_ID)
                .build();
        startForeground(256, notification);
    }

    @Nullable
    @Override
    public IBinder onBind(Intent intent) {
        Log.i("ServerManagerService", "onBind");
        return new RemoteServiceProxy();
    }

    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        Log.i("ServerManagerService", "onStartCommand");
        return super.onStartCommand(intent, flags, startId);
    }

    @Override
    public void onDestroy() {
        Log.i("ServerManagerService", "onDestroy");
        super.onDestroy();
    }
}
