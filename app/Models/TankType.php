<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class TankType extends Model
{

    public const TABLE_NAME = 'tank_types';

    public const WIDTH = 'width';
    public const HEIGHT = 'height';
    public const DEPTH = 'depth';
    public const GLASS_THICKNESS = 'glass_thickness';

    public $timestamps = false;
}
