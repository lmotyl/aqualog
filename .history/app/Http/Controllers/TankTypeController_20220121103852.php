<?php

namespace App\Http\Controllers;

use App\Models\TankType;

use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class TankTypeController extends Controller
{
    //

    public function getAll() 
    {
        $tankTypes = TankType::get()->toJson(JSON_PRETTY_PRINT);

        return response($tankTypes, Response::HTTP_OK);
    }

}
