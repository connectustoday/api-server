import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

export default class IAttachment {

	@prop({ required: true })
	schema_version: number;

	@prop({ required: true })
	type: string;

	@prop({ required: true })
	url: string;

	@prop()
	description?: string;

}
