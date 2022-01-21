<?php

namespace App\Http\Controllers;

use App\Models\Tank;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class TankController extends Controller
{
    //
    
    public function tanks()
    {
        $tanks = Tank::with(['user', 'tankType'])->get()->toJson(JSON_PRETTY_PRINT);

        return response($tanks, Response::HTTP_OK);
    }

}
