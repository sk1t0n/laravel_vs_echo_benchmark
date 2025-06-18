<?php

use App\Http\Controllers\TodoController;
use Illuminate\Foundation\Http\Middleware\VerifyCsrfToken;
use Illuminate\Support\Facades\Route;

Route::view('/', 'welcome');

Route::resource('/todos', TodoController::class)
    ->withoutMiddleware(VerifyCsrfToken::class);
