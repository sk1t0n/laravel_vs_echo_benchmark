<?php

namespace App\Http\Controllers;

use App\Models\Todo;
use Illuminate\Http\Request;

class TodoController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $todos = Todo::paginate(10);

        return response()->json([
            'todos' => $todos,
        ]);
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        $validated = $request->validate([
            'title' => 'required|max:255',
        ]);
        $todoCreated = Todo::create($validated);

        return response()->json([
            'todoCreated' => $todoCreated,
        ], 201);
    }

    /**
     * Display the specified resource.
     */
    public function show(Todo $todo)
    {
        return response()->json([
            'todo' => $todo,
        ]);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, Todo $todo)
    {
        $validated = $request->validate([
            'title' => 'required|max:255',
        ]);
        $ok = $todo->update($validated);

        return response()->json(['ok' => $ok]);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Todo $todo)
    {
        $ok = $todo->delete();

        return response()->json(['ok' => $ok]);
    }
}
