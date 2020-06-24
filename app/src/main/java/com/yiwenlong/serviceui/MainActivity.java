package com.yiwenlong.serviceui;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;

import com.yiwenlong.service.ServerManagerService;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        findViewById(R.id.start).setOnClickListener(view -> {
            Intent i = new Intent(getApplicationContext(), ServerManagerService.class);
            startForegroundService(i);
        });

        findViewById(R.id.stop).setOnClickListener(view -> {
            Intent i = new Intent(getApplicationContext(), ServerManagerService.class);
            stopService(i);
        });
    }
}